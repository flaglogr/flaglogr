<script setup lang="ts">
    import {
        useTheme
    } from 'vuetify';

    const theme = useTheme();
</script>

<template>
    <v-navigation-drawer
        expand-on-hover
        permanent
        rail
    >
        <v-list nav>
            <v-list-item
                prepend-avatar="https://randomuser.me/api/portraits/women/84.jpg"
                subtitle="jane.doe@example.ccom"
                title="Jane Doe"
                link
                href="/profile"
            />
        </v-list>

        <v-divider />

        <v-list density="compact" nav>
            <v-list-item
                prepend-icon="mdi-home"
                title="Home"
                value="home"
                link
                href="/"
            />
            <v-list-item
                prepend-icon="mdi-flag"
                title="Feature Flags"
                value="feature-flags"
                link
                href="/flags"
            />
            <v-list-item
                prepend-icon="mdi-ab-testing"
                title="A/B Testing"
                value="ab-testing"
                link
                href="/ab-testing"
            />
        </v-list>
        <template #append>
            <v-list density="compact">
                <v-list-item
                    prepend-icon="mdi-cog"
                    title="Settings"
                    value="settings"
                    link
                    href="/settings"
                />
                <v-list-item
                    prepend-icon="mdi-theme-light-dark"
                    title="Theme"
                    @click="theme.toggle()"
                />
            </v-list>
        </template>
    </v-navigation-drawer>
</template>
