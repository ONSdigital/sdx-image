# sdx-image

The sdx-image microservice is a bespoke app used within the Office for National Statistics (ONS) to aid with  
Survey Data eXchange (SDX). It provides functionality to transform a JSON survey response into a JPEG image.
This allows users downstream to process submissions collected through EQs (Electronic Questionnaires) with the same
set of tools as if the submission had been received in paper form and scanned in.


## Getting started
Ensure you have GO installed on your machine.

Clone the repo:
```bash
git clone git@github.com:ONSdigital/sdx-image.git
```

Navigate to the root of the project:
```bash
git cd sdx-image
```

Install dependencies:
```bash
go mod download
```

Build the application:
```bash
go build -o sdx-image
```

Run the executable
```bash
./sdx-image
```

Or, run the tests:
```bash
go test -json ./...
```
