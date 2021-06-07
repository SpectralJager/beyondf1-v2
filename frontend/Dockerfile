FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.16.0-alpine
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
#FROM nginx:1.18.0-alpine as production-stage
#RUN mkdir /app
#COPY --from=build-stage /app/dist /app
#COPY nginx.conf /etc/nginx/nginx.conf
#EXPOSE 8080
#CMD [ "npm", "run", "serve"]
#CMD ["nginx", "-g", "daemon off;"]
