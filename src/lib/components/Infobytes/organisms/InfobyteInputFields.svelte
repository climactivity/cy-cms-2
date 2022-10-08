<script lang="ts">
  import InfobyteCombobox from "../atoms/InfobyteCombobox.svelte";
  import InfobyteInput from "../atoms/InfobyteInput.svelte";
  import AspectAsyncCombobox from "../atoms/AspectAsyncCombobox.svelte";
  import FactorAsyncCombobox from "../atoms/FactorAsyncCombobox.svelte";
  import InfobyteValueCombobox from "../atoms/InfobyteValueCombobox.svelte";

  export let region: string;
  export let language: string;
  export let name: string;
  export let aspect: string;
  export let factor: string;
  export let difficulty: string;
  export let position: string;

  export let aspectsPromise: Promise<[string, number?][]>;
  export let factorsPromise: Promise<[string, number?][]>;
</script>

<div>
  <Col md="6" lg="4">
    <AspectAsyncCombobox
      inputId={"aspect"}
      name={"Aspekt"}
      bind:value={aspect}
      optionsPromise={aspectsPromise}
    />
  </Col>
  <Col md="6" lg="4">
    <FactorAsyncCombobox
      inputId={"factor"}
      name={"Gesichtspunkt"}
      bind:value={factor}
      optionsPromise={factorsPromise}
    />
  </Col>
  <Col>
    <InfobyteInput
      required={true}
      inputId={"position"}
      name={"Position"}
      bind:value={position}
    />
  </Col>
</div>

<div>
  <Col lg="12">
    <InfobyteInput
      required={true}
      inputId={"name"}
      name={"Name"}
      bind:value={name}
    />
  </Col>
</div>

<div>
  <Col md="6">
    <InfobyteCombobox
      inputId={"region"}
      name={"Region"}
      bind:value={region}
      options={["DE"]}
    />
  </Col>
  <Col md="6">
    <InfobyteCombobox
      inputId={"language"}
      name={"Sprache"}
      bind:value={language}
      options={["DE", "EN"]}
    />
  </Col>

  <Col md="6" lg="4">
    <InfobyteValueCombobox
      inputId={"difficulty"}
      name={"Schwierigkeit"}
      bind:value={difficulty}
      options={[
        ["Leicht", 1],
        ["Mittel", 2],
        ["Schwer", 3],
      ]}
    />
  </Col>
</div>
