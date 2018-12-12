# API

This document defines the API exposed by the nakolesach.sk.

## Create a Ride

```
POST https://us-central1-nakolesach-sk.cloudfunctions.net/createRide
```

## Close a Ride (before it's full)

```
POST https://us-central1-nakolesach-sk.cloudfunctions.net/closeRide
```


## Request a Ride

```
POST https://us-central1-nakolesach-sk.cloudfunctions.net/requestRide
```

## Accept a Request

```
POST https://us-central1-nakolesach-sk.cloudfunctions.net/acceptRequest
```

## Decline a Request

```
POST https://us-central1-nakolesach-sk.cloudfunctions.net/declineRequest
```

## Cancel a Request

```
DELETE https://us-central1-nakolesach-sk.cloudfunctions.net/cancelRequest
```