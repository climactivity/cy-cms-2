<script lang="ts">
    import { onMount } from "svelte";
    import type { NkStats, Tutorialcompletion } from "./AnalyticsTypes";
    import { baseUrl, isProd } from "$lib/stores/stores";
    import * as d3 from "d3";
    import moment from 'moment';
    let target: string;
    let stats: NkStats;

    let newInstallsDiv;
    let tutorialCompletionDiv;

    onMount(async () => {
        const targetData = await (
            await fetch(`${baseUrl}config`, {
                credentials: "include",
            })
        ).json();
        console.log(targetData);
        target = `${targetData.nk_server_host}/v2/rpc/get_stats?http_key=${targetData.nk_server_key}&unwrap`;

        const response = await fetch(target, {});
        stats = await (response.json() as Promise<NkStats>);
        console.log(stats);
        d3BarChart(newInstallsDiv,
            stats.new_installs.map(({count, date}) => {
             return {label: date.split('T')[0], value: count}
            }),
            [0,Math.max.apply(Math, stats.new_installs.map(({count, date}) => count))],
            newInstallsDiv.getBoundingClientRect().width, newInstallsDiv.getBoundingClientRect().height, "#77B924");

        d3BarChartH(tutorialCompletionDiv,
            stats.tutorial_completion.map(({count, key}) => {
             return {label: mapScreenNameToReadable(key), value: count}
            }).sort((a,b) => (a.value < b.value)),
            [0,Math.max.apply(Math, stats.tutorial_completion.map(({count, key}) => count))],
            tutorialCompletionDiv.getBoundingClientRect().width, tutorialCompletionDiv.getBoundingClientRect().height, "#64A6E2");
        
        
    });

    const d3BarChartH = (element, data: {label: string, value: number}[], yDomain, cWidth, cHeight, color) => {
        // set the dimensions and margins of the graph
        var margin = { top: 30, right: 30, bottom: 70, left: 120 },
            width = cWidth - margin.left - margin.right,
            height = cHeight - margin.top - margin.bottom;

        // append the svg object to the body of the page
        var svg = d3
            .select(element)
            .append("svg")
            .attr("width", width + margin.left + margin.right)
            .attr("height", height + margin.top + margin.bottom)
            .append("g")
            .attr(
                "transform",
                "translate(" + margin.left + "," + margin.top + ")"
            );

        // X axis
        const x = d3.scaleLinear().domain(yDomain).range([0, width]);
        svg
            .append("g")  
            .attr("transform", `translate(0, ${height})`)
            .call(d3.axisBottom(x));

        // Add Y axis
        const y = d3
            .scaleBand()
            .range([0, height])
            .domain(data.map((d) => d.label))
            .padding(0.2);
        svg.append("g")
            .call(d3.axisLeft(y))
            .selectAll("text")
            .style("text-anchor", "end");


        // Bars
        svg.selectAll("mybar")
            .data(data)
            .join("rect")
            .attr("x", 0)
            .attr("y", (d) => y(d.label))
            .attr("height", y.bandwidth())
            .attr("width", (d) => x(d.value))
            .attr("fill", color);
        
    };

    const d3BarChart = (element, data: {label: string, value: number}[], yDomain, cWidth, cHeight, color) => {
        // set the dimensions and margins of the graph
        var margin = { top: 30, right: 30, bottom: 70, left: 60 },
            width = cWidth - margin.left - margin.right,
            height = cHeight - margin.top - margin.bottom;

        // append the svg object to the body of the page
        var svg = d3
            .select(element)
            .append("svg")
            .attr("width", width + margin.left + margin.right)
            .attr("height", height + margin.top + margin.bottom)
            .append("g")
            .attr(
                "transform",
                "translate(" + margin.left + "," + margin.top + ")"
            );

        // X axis
        const x = d3
            .scaleBand()
            .range([0, width])
            .domain(data.map((d) => d.label))
            .padding(0.2);
        svg.append("g")
            .attr("transform", `translate(0, ${height})`)
            .call(d3.axisBottom(x))
            .selectAll("text")
            .style("text-anchor", "center");

        // Add Y axis
        const y = d3.scaleLinear().domain(yDomain).range([height, 0]);
        svg.append("g").call(d3.axisLeft(y));

        // Bars
        svg.selectAll("mybar")
            .data(data)
            .join("rect")
            .attr("x", (d) => x(d.label))
            .attr("y", (d) => y(d.value))
            .attr("width", x.bandwidth())
            .attr("height", (d) => height - y(d.value))
            .attr("fill", color);
        
    };

    const mapScreenNameToReadable = (screenName) => {
        switch (screenName) {
            case "AspectScreenPlayed":
                return "Aspekt angesehen"
            case "BigpointScreenPlayed":
                return "Zelt betreten"
            case "FirstInfobitPlayed":
                return "Ein Infobit angesehen"
            case "FirstQuestionPlayed":
                return "Eine Frage gespielt"
            case "GesichtspunktPlayed":
                return "Gesichtspunkt betreten"
            case "IntroPlayed":
                return "Wald betreten"
            case "QuestLogPlayed":
                return "Aspect Schirm"
            case "QuestPlayed":
                return "Eine Aufgabe angesehen"
            case "QuizIntroPlayed":
                return "Ein Infobyte betreten"
            case "ShopPlayed":
                return "Setzling ausgewählt"
            case "TrackingSettingsPlayed":
                return "Aspekt getrackt"
            case "WaterCollectPlayed":
                return "Wasser gesammelt"
            case "v2":
                return "App gestartet"
            defualt: 
                return screenName
        }
    }
</script>

<h1>App Statistiken</h1>
<section class="stats">
    <div class="card">
        <span class="title"> Installationen </span>
        <span class="value"> {#if stats} {stats.unique_installs} {/if}</span>
    </div>

    <div class="card">
        <span class="title"> Gepflanzte Bäume </span>
        <span class="value"> {#if stats} {stats.planted_trees} {/if}</span>
    </div>
    <div class="card">
        <span class="title"> Gespielte Quizze </span>
        <span class="value"> {#if stats} {stats.played_quizzes} {/if}</span>
    </div>
    <div class="card">
        <span class="title"> Abgeschlossene Challenges </span>
        <span class="value"> {#if stats} {stats.played_quests} {/if}</span>
    </div>

    <div bind:this={newInstallsDiv} class="chart card">
        <span class="title">Neuinstallationen über die letzen 7 Tage</span>
    </div>
    <div bind:this={tutorialCompletionDiv} class="chart card">
        <span class="title">Besuchte Bildschirme</span>
    </div>
</section>

<style>
    .stats {
        height: auto;
        width: 85vw;
        margin: 1rem auto;
        display: grid;
        gap: 1rem;
        grid-template-areas: 
            'unique_installs planted_trees played_quizzes played_quests '
            'new_installs new_installs new_installs new_installs'
            'tutorial_completion tutorial_completion tutorial_completion tutorial_completion';
        grid-template-rows: 5rem 800px 800px;
        grid-template-columns: repeat(4, 25%);
    }
    .card:nth-child(1) {
        grid-area: unique_installs;
    }
    .card:nth-child(2) {
        grid-area: planted_trees;
    }
    .card:nth-child(3) {
        grid-area: played_quizzes;
    }
    .card:nth-child(4) {
        grid-area: played_quests;
    }
    .card:nth-child(5) {
        grid-area: new_installs;
    }
    .card:nth-child(6) {
        grid-area: tutorial_completion;
    }
    .card {
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        align-items: center;
    }
    .card > span {
        margin:0 auto;
    }
    h1 {
        width: 100vw;
        margin:1rem auto;
        font-weight: lighter;
        text-align: center;
    }
    .title {
        font-weight: lighter;
    }
    .value {
        font-weight: bold;
        font-size: larger;
    }
</style>
