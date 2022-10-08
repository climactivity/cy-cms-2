export interface NkStats {
    planted_trees: number;
    played_quizzes: number;
    played_quests: number;
    unique_installs: number;
    new_installs: Newinstall[];
    tutorial_completion: Tutorialcompletion[];
}

export interface Tutorialcompletion {
    count: number;
    key: string;
}

export interface Newinstall {
    count: number;
    date: string;
}