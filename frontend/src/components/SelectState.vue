<script setup lang="ts">
import {ref, onMounted} from 'vue'
import {Loader} from '@googlemaps/js-api-loader';
import {useLazyQuery} from '@vue/apollo-composable'
import gql from 'graphql-tag'

interface State {
  name: string
}

const center = {lat: 37.1, lng: -95.7};

const selectedState = ref('')
const states = ref<State[]>([])

const querySearch = async (queryString: string, cb: (arg: any) => void) => {
  let url;
  if (queryString) {
    const queryState = encodeURIComponent(`{states(search: "${queryString}") {name}}`);
    url = `graphql?query=${queryState}`
  } else {
    const queryAll = encodeURIComponent(`{states{name}}`);
    url = `graphql?query=${queryAll}`
  }
  try {
    fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
        .then(res => res.json())
        .then(data => {
          states.value = data.data.states
          cb(states.value)
        })
  } catch (error) {
    console.error('Error fetching data:', error)
  }
}
const setCenter = () => {
  map.value = new google.maps.Map(mapContainer.value as HTMLDivElement, {
    zoom: 4,
    center: center
  })
}
const handleSelect = (item: Record<string, any>) => {
  const kml = encodeURIComponent(item.name);
  const url = `https://raw.githubusercontent.com/BattleWoLFz99/kml/refs/heads/main/${kml}.kml`
  const georssLayer = new google.maps.KmlLayer({
    url: url,
    zIndex: 10
  });
  setCenter()
  georssLayer.setMap(map.value? map.value: null);
}

const clearSelect = () => {
  setCenter()
}

const loader = new Loader({
  apiKey: "YOUR_API_KEY",
  version: "weekly",
  libraries: ["core", "maps", "places"]
});

const map = ref<google.maps.Map>()
const mapContainer = ref<HTMLDivElement>()
const markers = ref<google.maps.Marker[]>([])

onMounted(async () => {
  await loader.load();
  setCenter()
})

</script>

<template>

  <div class="container">
    <div class="input-container">
      <h1>United States Search</h1>
      <el-autocomplete
          value-key="name"
          v-model="selectedState"
          :fetch-suggestions="querySearch"
          clearable
          placeholder="Please input a state"
          @select="handleSelect"
          @clear="clearSelect"
      />
    </div>
    <div>
      <div ref="mapContainer" class="map"></div>
    </div>

  </div>
</template>

<style scoped>
.container {
  //display: flex;
  margin-top: 5%;
  margin-bottom: auto;
  width: 80%;
  height: 80%;
  //flex-direction: column;
}

.input-container {
  width: 20%;
  margin: auto;
  padding-bottom: 20px;
}

.map {
  width: 100%;
  height: 600px;
}
</style>