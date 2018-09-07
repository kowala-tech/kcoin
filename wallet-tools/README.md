## Kowala Wallet Tools

Install: `npm install`

Build (development): `npm run-script build`

Build (production): `npm run-script dist`

Note: The app's canonical URL and CDN URL are read from the environment (`APP_URL` and `CDN_URL`). They default to `https://wallet-tools.kowala.tech` and `https://cdn.kowala.tech`.  This is accomplished through templates via gulp.



## Building using docker for development

Build npm tools image and run npm install scripts:

```
docker build -t kowalatech/npm-tools -f wallet-tools/Dockerfile.npm .
docker run --rm -v $PWD/wallet-tools:/wallet-tools  kowalatech/npm-tools:latest npm install
```

Running the webserver
```
 docker build -t kowalatech/wallet-tools-nginx -f wallet-tools/Dockerfile.nginx .
 docker run --rm -d -v $PWD/wallet-tools/dist:/usr/share/nginx/html -p 443:443 kowalatech/wallet-tools-nginx:latest
```

And every time you make a change on the code execute to update the build.
```
docker run --rm -v $PWD/wallet-tools:/wallet-tools kowalatech/npm-tools:latest npm run-script build
```

As an alternative of the last command there is an option to watch the files that change and autobuild, on some systems it goes slowly. Try it at your own risk!
```
docker run --rm -v $PWD/wallet-tools:/wallet-tools kowalatech/npm-tools:latest gulp
```

#### Kowala Wallet Tools & Kowala Wallet Tools CX are licensed under The MIT License (MIT).
