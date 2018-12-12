# Models

This documentation explains the data model of nakolesach.sk. The model is used for data storage in
Firebase Firestore (thus some of entities are denormalized).

## General Information

All entities are looked up and manipulated solely by their keys. Every model thus has a unique
string ID. References are then made using the pattern `rideID: "Akfk9KoHUQ"` where the `ride`
refers to the `Ride` model. The same applies for normalized arrays, e.g. `rides: ["Akfk9KoHUQ"]`,
where the plural refers to a list of IDs of the `Ride` model.

## Ride Model

```json
{
  "id": "Akfk9KoHUQ",
  "createdAt": "0001-01-01T00:00:00Z",
  "rider": {
    "id": "4rIMp9HZAJ",
    "name": "Peter Malina"
  },
  
  "from": "Piestany",
  "to": "Brno",
  "slots": 4,
  
  "price": "6",
  "currency": "EUR",
  
  "open": true,
  "autoAccept": true,
  
  "boardingTime": "9999-12-31T23:59:59.999999999Z",
  "boardingPlace": "Shell, Piestany",
  "deboardingPlace": "Purkynova 127, Brno"
}
```

Where:
- `createdAt` is a time when the ride was created by its owner.
- `rider` is an embedded JSON with first part as user `id`, while the second part encodes users name.
- `from` and `to` are fully qualified addresses with at least first component (City)
- `slots` is a number of free slots in the car (required, does not default)
- `price` is the price for a single slot and `currency` is its currency (defaults to `EUR`)
- `open` indicates if the ride is open for requests.
- `autoAccept` will automatically accept any requests if set to `true`. All requests will need to be reviewed by the ride owner otherwise.
- `boardingTime` is the time of boarding for the ride.
- `boardingPlace` is a fully qualified address with all components. If specified, this is the boarding place, otherwise, the rider gets everybody from home.
- `deboardingPlace` is the same as boarding place, but with the exit specified instead.

Cases:
- (User) Find open rides from X to Y (`boardingTime > now() && open == true && from == X && to == Y`)

## Request Model

```json
{
  "id": "vm2uJ1N6JE",
  "createdAt": "0001-01-01T00:00:00Z",
  
  "slots": 2,
  
  "rideID": "Akfk9KoHUQ",
  "userID": "CUq3KOkb1V",
  "status": "OPEN"
}
```

Where:
- `slots` is a number of slots the user needs (default: 1).

Statuses Enum:
- `OPEN` means the ride was requested but not reviewed
- `APPROVED` means the request was accepted
- `DECLINED` means the request was declined
- `CANCELED` means the participant canceled the request (e.g. was boarded elsewhere)

Cases:
- (Driver) List open requests to my ride X (`rideID == X && status == 'OPEN'`)