<template>
    <div style="text-align: center;">
        <div class="view-buttons">
            <button @click="setGridView" :class="{ active: isGridView }">
                <img src="@/assets/images/icon/grid-view.svg" alt="Grid View">
            </button>
            <button @click="setListView" :class="{ active: !isGridView }">
                <img src="@/assets/images/icon/list-view.svg" alt="List View">
            </button>
        </div>
        <div v-if="!digimons">
            <p class="no-data">No data to display</p>
        </div>
        <div v-else>
            <div v-if="isGridView" class="grid-container" style="margin-top: 20px">
                <div v-for="digimon in digimons" :key="digimon.number" class="grid-item">
                    <img :src="`http://localhost:9706/${digimon.image}`" class="image" :alt="digimon.name"
                        :title="digimon.name">
                    <div class="flex-container">
                        <img class="element-grid" :alt="`${digimon.type}-${digimon.attribute}`"
                            :title="`${digimon.type}-${digimon.attribute}`"
                            :src="`http://localhost:9706/images/element/${toLowerCase(digimon.type)}-${toLowerCase(digimon.attribute)}.png`" />
                        <p class="name" style="margin-left:5px;">{{ digimon.name }}</p>
                    </div>
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
                    <div class="flex-container">
                        <div>
                            <img :src="`http://localhost:9706/${digimon.icon}`" class="icon" :alt="digimon.name"
                                :title="digimon.name">
                        </div>
                        <div style="margin-left: 10px;">
                            <div class="flex-container">
                                <img class="element-list" :alt="`${digimon.type}-${digimon.attribute}`"
                                    :title="`${digimon.type}-${digimon.attribute}`"
                                    :src="`http://localhost:9706/images/element/${toLowerCase(digimon.type)}-${toLowerCase(digimon.attribute)}.png`" />
                                <p class="name" style="font-size: 1em; margin-left:5px">{{ digimon.name }}</p>
                            </div>
                            <p class="number" style="margin-right:5px; font-size: 0.7em">
                                No. {{ formattedNumber(digimon.number) }}
                            </p>
                            <div class="stage-container">
                                <div :class="['stage-capsule', getStageClass(digimon.stage)]"
                                    style="font-size: 0.7rem;">
                                    {{ digimon.stage }}
                                </div>
                            </div>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
    </div>
    <div class="pagination-controls">
        <div class="page-buttons">
            <button @click="prevPage" :disabled="digimonListRequest.page_num === 1" class="arrow-button">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                    stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M15 18l-6-6 6-6" />
                </svg>
                <span class="sr-only">Previous</span>
            </button>
            <div class="page-numbers">
                <button v-for="page in totalPages" :key="page" @click="goToPage(page)"
                    :class="{ active: page === digimonListRequest.page_num }">
                    {{ page }}
                </button>
            </div>
            <button @click="nextPage" :disabled="digimonListRequest.page_num === totalPages" class="arrow-button">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                    stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M9 18l6-6-6-6" />
                </svg>
                <span class="sr-only">Next</span>
            </button>
        </div>
        <label style="font-weight: bold;" for="page-size">Page Size:</label>
        <select id="page-size" class="page-size" v-model="digimonListRequest.page_size" @change="fetchDigimons">
            <option v-for="size in pageSizes" :key="size" :value="size">{{ size }}</option>
        </select>
    </div>
</template>

<script>
import axios from '../../axios';

export default {
    props: {
        filters: {
            type: Object,
            default: () => ({
                stages: [],
                types: [],
                attributes: []
            })
        }
    },
    computed: {
        totalPages() {
            return Math.ceil(this.totalItems / this.digimonListRequest.page_size);
        }
    },
    data() {
        return {
            digimons: [],
            digimonListRequest: {
                page_size: 20,
                page_num: 1,
                sort_by: "number",
                sort_order: "asc"
            },
            pageSizes: [20, 40, 80],
            totalItems: 0,
            isGridView: true
        };
    },
    watch: {
        filters: 'resetPageAndFetch'
    },
    mounted() {
        this.fetchDigimons();
        this.fetchTotalItems();
    },
    methods: {
        async fetchDigimons() {
            const { stages, types, attributes } = this.filters;
            const requestBody = {
                ...this.digimonListRequest,
                stage: stages,
                type: types,
                attribute: attributes
            };
            try {
                const response = await axios.post('/digimon/list', requestBody, {
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });
                this.digimons = response.data
            } catch (error) {
                console.error('Error fetching Digimon data:', error);
            }
        },
        async fetchTotalItems() {
            const { stages, types, attributes } = this.filters;
            const requestBody = {
                ...this.digimonListRequest,
                stage: stages,
                type: types,
                attribute: attributes
            };
            try {
                const response = await axios.post('/digimon/list/size', requestBody, {
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });
                this.totalItems = response.data.size;
            } catch (error) {
                console.error('Error fetching total items:', error);
            }
        },
        updatePageSize() {
            this.digimonListRequest.page_num = 1;
            this.fetchDigimon();
            this.fetchTotalItems();
        },
        prevPage() {
            if (this.digimonListRequest.page_num > 1) {
                this.digimonListRequest.page_num--;
                this.fetchDigimons();
            }
        },
        nextPage() {
            this.digimonListRequest.page_num++;
            this.fetchDigimons();
        },
        goToPage(page) {
            if (page >= 1 && page <= this.totalPages) {
                this.digimonListRequest.page_num = page;
                this.fetchDigimons();
            }
        },
        resetPageAndFetch() {
            this.digimonListRequest.page_num = 1;
            this.fetchDigimons();
            this.fetchTotalItems();
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
        },
        toLowerCase(str) {
            return str.toLowerCase();
        }
    }
};
</script>

<style scoped>
.image {
    width: 150px;
}

.icon {
    width: 72px;
}

.element-grid {
    width: 30px;
    height: 30px;
}

.element-list {
    width: 20px;
    height: 20px;
}

.flex-container {
    display: flex;
}

.grid-container {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 10px;
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

.no-data {
    font-weight: bold;
    font-size: 2em;
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
    margin: 5px 0 0 0;
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
    border: 1px solid #000;
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

.pagination-controls {
    margin: 20px 0;
    display: flex;
    justify-content: flex-end;
    align-items: center;
}

.pagination-controls label {
    margin-right: 10px;
}

.pagination-controls select,
.pagination-controls input {
    margin-right: 20px;
}

.arrow-button {
    background: transparent;
    border: none;
    cursor: pointer;
    padding: 5px;
    display: flex;
    align-items: center;
}

.arrow-button svg {
    width: 24px;
    height: 24px;
}

.arrow-button:disabled {
    cursor: not-allowed;
    opacity: 0.5;
}

.sr-only {
    position: absolute;
    width: 1px;
    height: 1px;
    margin: -1px;
    padding: 0;
    overflow: hidden;
    clip: rect(0, 0, 0, 0);
    border: 0;
}

.page-buttons {
    margin-right: 10px;
    display: flex;
}

.page-size {
    margin-right: 0px;
    font-size: 0.8em;
    font-family: Avenir, Helvetica, Arial, sans-serif;
    font-weight: bold;
}

.page-numbers {
    display: flex;
    gap: 5px;
}

.page-numbers button {
    background: transparent;
    border: 0px;
    padding: 5px 10px;
    cursor: pointer;
    font-family: Avenir, Helvetica, Arial, sans-serif;
    font-weight: bold;
}

.page-numbers button.active {
    background-color: #374e98;
    color: #fff;
}
</style>