import { baseUrl, Infobyte } from '$lib/stores/stores'
import type { infobyteIdentifier } from "./particles/Infobyte";

export const getInfobyte = async (id: string): Promise<Infobyte> => {
    if (!id) return;

    const infobyte = fetch(`${baseUrl}infobyte/${id}`, { credentials: "include", })
        .then(data => data.json())
        .then(json => json as Infobyte);

    return infobyte;
};

export const getInfobytes = async (): Promise<infobyteIdentifier[]> => {

    const response = fetch(`${baseUrl}infobyte`, { credentials: "include", })
        .then(data => data.json())
        .then(json => json as infobyteIdentifier[])

    return response;
};

export const fetchAspects = async (region: string, language: string): Promise<any[]> => {

    const response = fetch(`${baseUrl}localized-aspect?${new URLSearchParams({ r: region, l: language, })}`
        , { credentials: "include", })
        .then(data => data.json());

    return response;
};

export const deleteInfobyte = async (id: string): Promise<any> => {

    const response = fetch(`${baseUrl}infobyte/${id}`, {
        credentials: "include",
        method: "DELETE",
    }).then(data => data.json());

    return response;
}

export const createOrUpdateInfobyte = async (infobyte: Infobyte): Promise<any> => {
    const response = fetch(`${baseUrl}infobyte${infobyte._id ? "/" + infobyte._id : ""}`,
        {
            credentials: "include",
            method: infobyte._id ? "PUT" : "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(infobyte),
        });
    return response;
}