import type PocketBase from 'pocketbase';


export const findRecord = (pocketBase: PocketBase, searchParam) => {
    const {baseUrl} = pocketBase

   return pocketBase.send(`find/${searchParam}`, {
        
   })
}