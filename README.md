# About

This is the repo for the new [PocketBase](https://pocketbase.io/) based CMS for our app. 

It's not very complete yet
## Developing

You'll need [go](https://go.dev/) and npm

Run 
``` bash
make all
```
 to 
  - fetch the js dependencies  with `npm install`
  - build the frontend with `npm run build`
  - and finally compile and run the go part with `go run ...`
    - go will fetch dependencies automatically on first build


### Hacking on the frontend
Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:


```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

### working on the backend
go sources are located in the `cms` directory, migrations go in the `migrations` dir.
### building a containerized version
run 

``` bash
make run-container
```
 to deploy a version with docker-compose.

## Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

The CMS frontend uses the static site adapter `@sveltejs/adpter-static` and uses it's build output as the public dir `pb_public` for PocketBase.

