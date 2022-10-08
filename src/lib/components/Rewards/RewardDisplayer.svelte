<script lang="ts">
import type { Reward } from "./RewardTypes";
export let reward: Reward;
</script>

<span class="reward">{`XP: ${reward.xp} T: ${reward.coins} W: ${reward.water} `}</span>

<style>
    .reward {
        border: 1px;
        border-radius: 4px;
        border-style: solid;
        padding: 2px;
    }
</style>