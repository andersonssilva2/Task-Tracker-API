FROM node:20-alpine AS build
WORKDIR /app
COPY package.json tsconfig.json vite.config.mjs ./
COPY index.html ./
COPY src ./src
COPY public ./public
RUN npm install && npm run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
