<script lang="ts">

import { baseUrl } from "$lib/stores/stores";

    export let data
</script>

<div>
    <img src="https://app.climactivity.de{data.relativeUrl}" class="thumb"/>
</div>

<style lang="scss">
.thumb {
    width: 200px;
    height: 200px;
    object-fit: contain;
}
</style>