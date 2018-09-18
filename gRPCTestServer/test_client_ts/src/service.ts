import { grpc } from "grpc-web-client";

import { GetMyCatMessage, MyCatResponse } from "./service_pb";
import { Cat, CatClient, ServiceError } from "./service_pb_service";

const postMessage = (value: string) => {
    const getMyCatMessage = new GetMyCatMessage();
    getMyCatMessage.setTargetCat(value);

    //const catClient = new CatClient("127.0.0.1:19003");
    grpc.invoke(Cat.GetMyCat, {
        host: "http://localhost:19003",
        request: getMyCatMessage,
        onMessage: (message) => {
            console.log("onMessage", message.toObject());
          },
        onEnd: (code, msg, trailers) => {
            console.log("onEnd", code, msg, trailers);
        },
    })
    /*
    catClient.getMyCat(getMyCatMessage, (error: ServiceError, response: MyCatResponse | null) => {
        console.log(error);
        console.log(response);
    });*/
}

postMessage("tama");
