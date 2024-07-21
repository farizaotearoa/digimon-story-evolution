<template>
    <div style="text-align: center;">
        <img :src="`http://localhost:9706/images/title/digimon-list.png`" class="title" alt="Digimon title">
        <div class="view-buttons">
            <button @click="setGridView" :class="{ active: isGridView }">
                <img src="@/assets/images/icon/grid-view.svg" alt="Grid View">
            </button>
            <button @click="setListView" :class="{ active: !isGridView }">
                <img src="@/assets/images/icon/list-view.svg" alt="List View">
            </button>
        </div>
        <div v-if="isGridView" class="grid-container" style="margin-top: 20px">
            <div v-for="digimon in digimons" :key="digimon.number" class="grid-item">
                <img :src="`http://localhost:9706/${digimon.image}`" class="image" alt="{{ digimon.name }}">
                <p class="name">{{ digimon.name }}</p>
                <p class="number">No. {{ formattedNumber(digimon.number) }}</p>
                <div class="stage-container">
                    <div :class="['stage-capsule', getStageClass(digimon.stage)]">
                        {{ digimon.stage }}
                    </div>
                </div>
            </div>
        </div>
        <ul v-else class="list-container">
            <li v-for="digimon in digimons" :key="digimon.number" class="list-item">
                <p class="name" style="font-size: 1em;">{{ digimon.name }}</p>
                <p class="number" style="margin-right:5px; font-size: 0.7em">
                    No. {{ formattedNumber(digimon.number) }}
                </p>
                <div class="stage-container">
                    <div :class="['stage-capsule', getStageClass(digimon.stage)]" style="font-size: 0.7rem;">
                        {{ digimon.stage }}
                    </div>
                </div>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            isGridView: true,
            digimons: []
        };
    },
    created() {
        this.fetchDigimons();
    },
    methods: {
        async fetchDigimons() {
            try {
                const response = await fetch('http://localhost:9706/digimon/list');
                this.digimons = await response.json();
            } catch (error) {
                console.error('Error fetching Digimon data:', error);
            }
        },
        getStageClass(stage) {
            switch (stage) {
                case 'Training 1':
                    return 'stage-training1';
                case 'Training 2':
                    return 'stage-training2';
                case 'Rookie':
                    return 'stage-rookie';
                case 'Champion':
                    return 'stage-champion';
                case 'Ultimate':
                    return 'stage-ultimate';
                case 'Mega':
                    return 'stage-mega';
                case 'Ultra':
                case 'Armor':
                    return 'stage-special';
                default:
                    return 'no-stage';
            }
        },
        setGridView() {
            this.isGridView = true;
        },
        setListView() {
            this.isGridView = false;
        },
        formattedNumber(number) {
            if (number < 10) {
                return `00${number}`;
            } else if (number < 100) {
                return `0${number}`;
            } else {
                return number.toString();
            }
        }
    }
};
</script>

<style scoped>
.title {
    width: 500px;
}

.image {
    width: 250px;
}

.grid-container {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 20px;
}

.grid-item {
    background-color: #fff;
    border-radius: 5px;
    padding: 20px;
    text-align: center;
    border: 2px solid #374e98;
}


.list-container {
    list-style-type: none;
    padding: 0;
}

.list-item {
    align-items: center;
    border: 2px solid #374e98;
    border-radius: 5px;
    padding: 20px;
    margin-bottom: 10px;
}

.name {
    font-style: normal;
    font-weight: bold;
    font-size: 1.5em;
    text-align: left;
    margin: 0;
}

.number {
    font-style: normal;
    font-weight: bolder;
    font-size: 1em;
    text-align: left;
    margin: 0;
}

.stage-container {
    display: flex;
    justify-content: flex-start;
    width: 100%;
}

.stage-capsule {
    color: #fff;
    padding: 5px 15px;
    border-radius: 50px;
    font-size: 0.9rem;
    margin-top: 10px;
}

.stage-training1 {
    background-color: #8ab4d9;
}

.stage-training2 {
    background-color: #274b96;
}

.stage-rookie {
    background-color: #29b376;
}

.stage-champion {
    background-color: #016f3c;
}

.stage-ultimate {
    background-color: #f9a73e;
}

.stage-mega {
    background-color: #bf212f;
}

.stage-special {
    background-color: #FFD700;
    color: #545454;
}

.no-stage {
    background-color: #fff;
    color: #000;
}

.subtitle {
    font-family: 'Digivolve', sans-serif;
    font-style: normal;
    font-size: 4rem;
    margin: 0px 10px;
    background: linear-gradient(to bottom, #fbda04, #fa700a);
    /* background: linear-gradient(to bottom, #fff, #fff); */
    color: transparent;
    position: relative;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
    display: inline-block;
}

.subtitle::before {
    content: 'Digimon List';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: -1;
    color: #374e98;
    font-family: 'Digivolve', sans-serif;
    font-size: 4rem;
    font-weight: bold;
    text-shadow:
        3px -1px 0 #374e98,
        0px -3px 0 #374e98,
        1px 2px 0 #374e98,
        3px 1px 0 #374e98;
    display: block;
}

.view-buttons {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
}

.view-buttons button {
    border: none;
    background: none;
    cursor: pointer;
    padding: 7px;
    padding-bottom: 4px;
    border-radius: 5px;
    transition: background-color 0.3s;
}

.view-buttons button img {
    width: 15px;
    height: 15px;
}

.view-buttons button.active {
    background-color: #f0f0f0;
}
</style>