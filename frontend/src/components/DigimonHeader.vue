<template>
    <div class="title">
        <img :src="`http://localhost:9706/images/title/digimon-title.png`" class="image" alt="Digimon title">
    </div>
</template>

<script>
export default {
    data() {
        return {};
    }
};
</script>

<style scoped>
.title {
    display: inline-block;
    padding: 10px;
}

.image {
    width: 250px;
}
</style>