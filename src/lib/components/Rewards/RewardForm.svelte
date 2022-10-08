<script lang="ts">
  import type { Reward } from "./RewardTypes";

  export let reward: Reward;
</script>

<div>
  <h5 class="mt-3">Reward</h5>

  <label for="xp">xp</label>
  <input id="xp" placeholder="xp" bind:value={reward.xp} />
  <label for="coins">coins</label>
  <input id="coins" placeholder="coins" bind:value={reward.coins} />
  <label for="water">water</label>
  <input id="water" placeholder="water" bind:value={reward.water} />
</div>
