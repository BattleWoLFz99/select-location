FROM guergeiro/pnpm:22-9-alpine AS base

COPY . /app
WORKDIR /app

RUN pnpm install --frozen-lockfile
RUN pnpm build

FROM nginx:1.27-alpine
COPY --from=0 /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]