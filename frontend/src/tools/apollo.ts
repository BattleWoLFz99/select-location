import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client/core'


const httpLink = createHttpLink({
    uri: 'http://localhost:8080/graphql',
    headers: {
        'Content-Type': 'application/json',
    },
})
const cache = new InMemoryCache()
const apolloClient = new ApolloClient({
    link: httpLink,
    cache,
})

export default apolloClient