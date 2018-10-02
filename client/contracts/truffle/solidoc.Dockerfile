FROM microsoft/dotnet:2.1-sdk-alpine AS build-env
RUN apk update && apk add git nodejs python g++ make bash && rm -rf /var/cache/apk/*

WORKDIR /app
RUN git clone https://github.com/binodnp/solidoc .
RUN dotnet restore
RUN dotnet publish -c Release

RUN npm install -g truffle@4.1.x

ENTRYPOINT ["dotnet", "/app/bin/Release/netcoreapp2.1/publish/Solidoc.dll", "/src", "/out"]