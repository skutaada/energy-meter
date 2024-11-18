FROM oven/bun:alpine AS builder

WORKDIR /app
COPY . .
ARG VITE_APP_BACKEND_ADDRESS
ENV VITE_APP_BACKEND_ADDRESS $VITE_APP_BACKEND_ADDRESS
RUN bun install
RUN bun run build

FROM nginx:latest-alpine-slim AS prod
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf  /etc/nginx/conf.d
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
