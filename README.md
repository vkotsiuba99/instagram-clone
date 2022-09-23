# Instagram Clone

> A scalable Instagram clone powered by microservices written in Golang

## Getting started

Setup frontend:

```bash
$ git clone https://github.com/vkotsiuba99/instagram-clone
$ cd instagram-clone/frontend
$ npm install
# or
$ yarn
```

Start the frontend:

```bash
$ npm run dev
# or
$ yarn dev
```

The project should be located on `http://localhost:3000`.

## Services

### Account API ([service.account-api](https://github.com/vkotsiuba99/instagram-clone/tree/master/service.account-api))

RESTful Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a account.

### Image Storage ([service.image-storage](https://github.com/vkotsiuba99/instagram-clone/tree/master/service.image-storage))

Go based image service supporting Gzipped content, multi-part forms and a RESTful approach for uploading and downloading images.

### Frontend ([frontend](https://github.com/vkotsiuba99/instagram-clone/tree/master/frontend))

React.js webapp that represents a refresh Instagram UI presenting different information from the services.

This project uses [Vite](http://vitejs.dev/) for a fast bundling and a better development experience.