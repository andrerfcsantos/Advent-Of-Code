import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import Home from "../views/Home.vue";
import Day01 from "../views/2019/Day01.vue";
import Day01_2020 from "../views/2020/Day01_2020.vue";
import NotFound from "../views/NotFound.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/2019/day/1",
    name: "2019_day01",
    component: Day01
  },
  {
    path: "/2020/day/1",
    name: "2020_day01",
    component: Day01_2020
  },
  {
    path: "*",
    name: "not_found",
    component: NotFound
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
