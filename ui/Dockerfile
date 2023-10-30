FROM node:latest as frontendbuild
WORKDIR /usr/local/app
COPY ./ /usr/local/app/
RUN npm install
RUN npm run build

FROM nginx:latest as final
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=frontendbuild /usr/local/app/dist/songbird /usr/share/nginx/html