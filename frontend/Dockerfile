FROM oven/bun:alpine AS builder

WORKDIR /app
COPY . .
ARG VITE_APP_BACKEND_ADDRESS
ENV VITE_APP_BACKEND_ADDRESS $VITE_APP_BACKEND_ADDRESS
RUN bun install
RUN bun run build
EXPOSE 80
CMD ["bunx", "serve", "-l", "tcp://0.0.0.0:80", "-C", "dist"]

