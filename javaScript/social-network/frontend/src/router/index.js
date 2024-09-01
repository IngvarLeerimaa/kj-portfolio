import { createRouter, createWebHashHistory } from "vue-router";
import WSConnection from "../assets/websocket.js";
import HomeView from "../views/HomeView.vue";
import LoginView from "../views/LoginView.vue";
import RegisterView from "../views/RegisterView.vue";
import app from "@/main.js";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
    meta: { requiresAuth: true },
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/register",
    name: "register",
    component: RegisterView,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth && document.cookie.split("sessionID=").length <= 1) {
    next("/login");
    return;
  }

  if (to.meta.requiresAuth && WSConnection.ws == null) {
    await fetch("http://localhost:3000/api/v1/user/session", {
      method: "GET",
      credentials: "include",
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Unauthorized");
        }
        return response.json();
      })
      .then((data) => {
        if (data.userid > 0) {
          app.config.globalProperties.$userId = data.userid;
          next();
        } else {
          next("/login");
        }
      })
      .catch((error) => {
        console.error(error);
        next("/login");
      });
  } else {
    next();
  }
});

export default router;
