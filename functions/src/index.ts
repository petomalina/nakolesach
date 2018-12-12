import * as functions from 'firebase-functions';
import {Request, Response} from "express";

export const createRide = functions.https.onRequest((req: Request, res: Response) => {
    res.send("Not implemented yet");
});

export const closeRide = functions.https.onRequest((req: Request, res: Response) => {
    res.send("Not implemented yet");
});

export const createRequest = functions.https.onRequest((req: Request, res: Response) => {
    res.send("Not implemented yet");
});

export const acceptRequest = functions.https.onRequest((req: Request, res: Response) => {
    res.send("Not implemented yet");
});

export const declineRequest = functions.https.onRequest((req: Request, res: Response) => {
    res.send("Not implemented yet");
});

export const cancelRequest = functions.https.onRequest((req: Request, res: Response) => {
    res.send("Not implemented yet");
});
