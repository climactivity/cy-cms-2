<script lang="ts">
  import ContentDebug from "../Editor/content-debug.svelte";

  import ContentEditorSection from "../Editor/content-editor-section.svelte";
  import ContentEditor from "../Editor/content-editor.svelte";
  import SelectEdit from "../Editor/select-edit.svelte";

  import StringEdit from "../Editor/string-edit.svelte";
  import type Aspect from "./AspectTypes";
  import FactorEdit from "./factor-edit.svelte";
  import TrackingEdit from "./tracking-edit.svelte";

  export let data: Aspect = {
    themenmonat: null,
    templateType: "tree",
    published: false,
    name: "",
    title: "",
    frontmatter: "",
    localizedTrackingData: {
      question: "",
      footnote: "",
      options: [],
    },
    forRegion: "DE",
    forLanguage: "DE",
    localizedFactors: [],
  };
</script>

<ContentEditor bind:data bind:title={data.title} saveTarget="localized-aspect">
  <ContentEditorSection label="Metadata">
    <!-- Name -->
    <StringEdit
      label="Name"
      type="string"
      placeholder="Name"
      bind:value={data.name}
      id="name"
    />

    <!-- Title -->
    <StringEdit
      label="Titel"
      type="string"
      placeholder="Titel der in der App angezeigt wird"
      bind:value={data.title}
      id="title"
    />

    <!-- Sektor -->
    <SelectEdit
      id="sector"
      label="Sektor"
      placeholder="-- Sektor auswählen --"
      bind:value={data.bigpoint}
      options={[
        { value: "ernaehrung", label: "Zelt der Ernährung" },
        { value: "energy", label: "Zelt der Energie" },
        { value: "mobility", label: "Zelt der Mobilität" },
        { value: "indirect_emissions", label: "Zelt des Konsums" },
        {
          value: "private_engagement",
          label: "Zelt des privaten Engagements",
        },
        {
          value: "public_engagement",
          label: "Zelt des öffentlichen Engagements",
        },
        { value: "accounting", label: "Bilanzierung" },
      ]}
    />

    <!-- Themenmonat -->
    <SelectEdit
      id="themenmonat"
      label="Themenmonat"
      placeholder="-- Themenmonat auswählen --"
      bind:value={data.themenmonat}
      options={[
        { value: "bilanzierung", label: "Bilanzierung" },
        { value: "ernaehrung", label: "Ernährung" },
        { value: "sich_informieren", label: "Sich informieren" },
        {
          value: "mobilitaet_zu_lande",
          label: "Mobilität zu Lande",
        },
        {
          value: "strom_produzieren_und_sparen",
          label: "Strom produzieren und sparen",
        },
        {
          value: "klimaschutzengagement_als_privatperson",
          label: "Klimaschutzengagement als Privatperson",
        },
        { value: "dinge_teilen", label: "Dinge teilen" },
        {
          value: "klimaschutz_und_geld",
          label: "Klimaschutz und Geld",
        },
        { value: "waerme", label: "Wärme" },
        {
          value: "klimaschutz_in_schule_und_betrieb",
          label: "Klimaschutz in Schule und Betrieb",
        },
        { value: "bewusst_einkaufen", label: "Bewusst einkaufen" },
        {
          value: "urlaub_und_fliegen",
          label: "Urlaub und Fliegen",
        },
        {
          value: "gemeinsam_fuer_den_klimaschutz",
          label: "Gemeinsam für den Klimaschutz",
        },
      ]}
    />

    <!-- Template Type (sets what type of thing is awarded in the app) -->
    <SelectEdit
      id="template-type"
      label="Template Type"
      placeholder="-- Entitätenart --"
      bind:value={data.templateType}
      options={[
        { value: "tree", label: "Baum" },
        { value: "bush", label: "Busch" },
        { value: "small", label: "Kleinkram" },
      ]}
    />

    <!-- Region, currently we're only in DE, but this will hopefully change -->
    <SelectEdit
      id="region"
      label="Region"
      placeholder="-- Region --"
      bind:value={data.forRegion}
      options={[{ value: "DE", label: "DE" }]}
    />

    <!-- Language -->
    <SelectEdit
      id="language"
      label="Sprache"
      placeholder="-- Sprache --"
      bind:value={data.forLanguage}
      options={[
        { value: "DE", label: "DE" },
        { value: "EN", label: "EN" },
      ]}
    />

    <!-- message: an introduction why this is important, shown when entering the aspect-->
    <StringEdit
      label="Einleitungstext"
      type="textarea"
      placeholder="Beschreibung"
      bind:value={data.frontmatter}
      id="message"
    />
  </ContentEditorSection>

  <ContentEditorSection label="Gesichtspunkte">
    <FactorEdit bind:value={data.localizedFactors} />
  </ContentEditorSection>

  <ContentEditorSection label="Tracking">
    <TrackingEdit bind:value={data.localizedTrackingData} />
  </ContentEditorSection>
  <ContentDebug {data} />
</ContentEditor>
