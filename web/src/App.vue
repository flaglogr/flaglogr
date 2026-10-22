<template>
    <v-app>
        <side-bar />
        <v-main>
            <router-view />
        </v-main>
    </v-app>
</template>

<script lang="ts" setup>
    import SideBar from './components/SideBar.vue';
</script>
