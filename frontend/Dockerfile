FROM oven/bun:alpine AS builder

WORKDIR /app

COPY package*.json bun.lockb ./
RUN bun install
COPY . .
ARG BACKEND_ADDRESS
ENV BACKEND_ADDRESS $BACKEND_ADDRESS
RUN bun run build

FROM nginx:stable-alpine AS production

COPY --from=builder /app/nginx /etc/nginx/conf.d
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 8080
ENTRYPOINT ["nginx", "-g", "daemon off;"] 

