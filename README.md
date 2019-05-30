# Seeds

<a href="https://cloud.drone.io/CloudHammer/Seeds"><img src="https://cloud.drone.io/api/badges/CloudHammer/Seeds/status.svg" alt="Build Status"></a>
<a href="https://microbadger.com/images/cloudhammer/seeds"><img src="https://images.microbadger.com/badges/image/cloudhammer/seeds.svg" alt="Get your own image badge on microbadger.com"></a>
<a href="https://opensource.org/licenses/GPL-3.0"><img src="https://img.shields.io/badge/license-GPL--3.0-blue.svg" alt="License"></a>

High performance mod multi user API implementation

## Quick start
### Using Docker
```
docker run -d \
    --name seeds \
    -e DATBASE_PASS=password \
    cloudhammer/seeds
```
### Build with go
```
cd $GOPATH/src
mkdir -p github.com/CloudHammer
cd github.com/CloudHammer
git clone https://github.com/CloudHammer/Seeds.git
cd ~
go build -o seeds github.com/CloudHammer/Seeds/src/main
```