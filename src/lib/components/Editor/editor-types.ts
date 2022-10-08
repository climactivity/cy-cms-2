import type { SvelteComponent } from "svelte"

export type ContentMetadata = {
    data: Content[],
    count: number
}

export type Content = {
    _id?: string, 
    published: boolean,
    [key:string]: any
}

export type FieldExtractor = {
    component: any,
    name:string,
    type: "string" | "RichText" | "number" | "select",
    label: string,
    placeholder?: string,
    range?: {min: number, max:number, step:number},

    selectOptions?: {
        key:string, value: any
    }
}

export type FieldSetExtractor = {
    [key:string]: FieldExtractor
}

