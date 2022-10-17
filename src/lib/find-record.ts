import type PocketBase from 'pocketbase';

export const findRecord = (pocketBase: PocketBase, searchParam) => {
	return pocketBase.send(`find/${searchParam}`, {});
};
