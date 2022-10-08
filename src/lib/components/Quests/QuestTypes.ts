import { Reward } from "../Rewards/RewardTypes";

export const enum FilterType {
    "REQURIED", "EXCLUDED", "PREFERED" 
}

export class QuestAspectFilter {
    filterType: FilterType = FilterType.EXCLUDED
    aspectId: string
    trackingLevel: Number
}

export class QuestQuestFilter {
    filterType: FilterType = FilterType.EXCLUDED
    questId: string
}

export class QuestDto {
    _id?: string;
    deadline: string;
    maxDuration: string;
    startDate: string;
    region: string; 
    language: string
    title: string
    text: any;
    baseReward: Reward;
    published: boolean = false;
    questAspectFilter: QuestAspectFilter[] = []; 
    questQuestFilter: QuestQuestFilter[] = [];
    questFollowup: string; // immediately following quest
    alertTrackedAspect: string; // affected aspect _id field
    linkToAfter: string; 
    altIcon: string;
    userSelectable: boolean = true;
    reward: Reward = new Reward();
    triggerTrackingUpdate: boolean = false;
    numRepeat: string = "1";
    repeatAfter: string = "0";

}
    