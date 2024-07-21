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

.image-container {
    background-color: black;
    /* background: linear-gradient(to bottom, #fbda04, #fa700a); */
    padding: 10px;
    display: inline-block;
    position: absolute;
    top: 5px;
    left: 5px;
}

.image-container img {
    display: block;
}

.image {
    width: 250px;
}
</style>