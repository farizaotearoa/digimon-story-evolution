<template>
    <div class="filter-component">
        <div class="filter-container">
            <p class="title">Filter</p>
            <div class="stage">
                <p class="filter-title">Stage</p>
                <div v-for="stage in stages" :key="stage">
                    <input type="checkbox" :id="'stage-' + stage" :value="stage" v-model="selectedStages" />
                    <label :for="'stage-' + stage">{{ stage }}</label>
                </div>
            </div>
            <div class="type">
                <p class="filter-title">Type</p>
                <div v-for="type in types" :key="type">
                    <input type="checkbox" :id="'type-' + type" :value="type" v-model="selectedTypes" />
                    <label :for="'type-' + type">{{ type }}</label>
                </div>
            </div>
            <div class="attribute">
                <p class="filter-title">Attribute</p>
                <div v-for="attribute in attributes" :key="attribute">
                    <input type="checkbox" :id="'attribute-' + attribute" :value="attribute"
                        v-model="selectedAttributes" />
                    <label :for="'attribute-' + attribute">{{ attribute }}</label>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
export default {
    data() {
        return {
            stages: ['Training 1', 'Training 2', 'Rookie', 'Champion', 'Ultimate', 'Mega', 'Ultra', 'Armor', 'No Stage'],
            selectedStages: [],
            types: ['Free', 'Vaccine', 'Virus', 'Data'],
            selectedTypes: [],
            attributes: ['Dark', 'Earth', 'Electric', 'Fire', 'Light', 'Neutral', 'Plant', 'Water', 'Wind'],
            selectedAttributes: []
        }
    },
    watch: {
        selectedStages: 'emitFilterChange',
        selectedTypes: 'emitFilterChange',
        selectedAttributes: 'emitFilterChange'
    }, methods: {
        emitFilterChange() {
            this.$emit('filter-change', {
                stages: this.selectedStages,
                types: this.selectedTypes,
                attributes: this.selectedAttributes
            });
        }
    }
}
</script>

<style scoped>
.filter-component {
    padding: 20px;
    padding-top: 0px;
    height: 100vh;
}

.filter-container {
    padding: 10px;
    text-align: left;
    border-radius: 5px;
    border: 2px solid #374e98;
}

.filter-container .stage,
.filter-container .type {
    margin-bottom: 10px;
}

.filter-container .stage,
.filter-container .type {
    padding-bottom: 7px;
    border-bottom: 1px solid #a9a9a9;
}

.title {
    margin-top: 0px;
    margin-bottom: 5px;
    font-weight: bold;
    font-size: 1.3em;
    text-align: center;
}

.filter-title {
    margin-top: 0px;
    margin-bottom: 8px;
    font-weight: bold;
    font-size: 1em;
    color: #757575;
}
</style>