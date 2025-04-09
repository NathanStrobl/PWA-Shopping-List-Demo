FROM node:23.10.0 AS build
WORKDIR /app
COPY . . 
RUN npm install
RUN npm run build

FROM node:23.10.0 AS prod
WORKDIR /app
COPY --from=build /app/dist
RUN npm install -g serve
EXPOSE 3000
CMD ["serve", "-s", ".", "-l", "3000"]