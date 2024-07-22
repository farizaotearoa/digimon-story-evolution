<template>
    <div class="layout">
        <img :src="`http://localhost:9706/images/title/digimon-list.png`" class="title" alt="Digimon title">
        <div class="main-layout">
            <DigimonListFilter class="filter-section" @filter-change="updateFilters" />
            <div class="list-section">
                <DigimonList :filters="filters" />
            </div>

        </div>
    </div>
</template>

<script>
import DigimonListFilter from './DigimonListFilter.vue'
import DigimonList from './DigimonList.vue'

export default {
    components: {
        DigimonListFilter,
        DigimonList
    },
    data() {
        return {
            filters: {
                stages: [],
                types: [],
                attributes: []
            }
        };
    },
    methods: {
        updateFilters(newFilters) {
            this.filters = {
                stages: newFilters.stages || [],
                types: newFilters.types || [],
                attributes: newFilters.attributes || []
            };
        }
    }
}
</script>

<style scoped>
.main-layout {
    display: flex;
    height: 100vh;
}

.layout {
    text-align: center;
    width: 1450px;
    margin-left: auto;
    margin-right: auto;
}

.filter-section {
    flex: 1 1 20%;
    background-color: #fff;
}

.list-section {
    flex: 1 1 80%;
    background-color: #fff;
}

.title {
    width: 500px;
    margin-bottom: 50px;
}
</style>