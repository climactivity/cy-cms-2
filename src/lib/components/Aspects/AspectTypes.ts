import type { Content } from "../Editor/editor-types";

export default interface Aspect extends Content {
    themenmonat: string,
    templateType: string,
    published: boolean;
    _id?: string;
    name: string;
    title: string;
    frontmatter: string;
    icon?: any;
    infoGraph?: any;
    localizedTrackingData?: LocalizedTrackingData;
    forRegion: string;
    forLanguage: string;
    message?: string;
    bigpoint?: string;
    localizedFactors?: LocalizedFactor[]

}

export interface LocalizedTrackingData {
    question: string;
    footnote: string;
    options: LocalizedTrackingOption[]
}

export interface LocalizedTrackingOption {
    option: string,
    reward: {xp: number, coins: number, water:number},
    level: number,
    co2value?: number,
    waterFactor: number,
    editing?: boolean
}

export interface LocalizedFactor {
        editing?: boolean;
        id: number,
        name: string,
        intro: string
}
export interface TrackingOption {
    id: string,
    answer: string,
    xp: number,
    coins: number,
    water: number
}