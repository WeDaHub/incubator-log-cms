import Vue from "vue";
import Router from "vue-router";
import HelloWorld from "@/components/HelloWorld";
import Login from "@/components/Login";
import Main from "@/components/Main";
import Log from "@/components/Log";
import ScanWeekLogs from "@/components/ScanWeekLogs";
import LogList from "@/components/LogList";

Vue.use(Router);

Vue.component("myLogList", LogList);
Vue.component("myLog", Log);

export default new Router({
    routes: [{
            path: "/",
            name: "Login",
            component: Login
        },
        {
            path: "/Main",
            name: "Main",
            component: Main
        },
        {
            path: "/ScanWeekLogs",
            name: "ScanWeekLogs",
            component: ScanWeekLogs
        },
        {
            path: "/Log",
            name: "Log",
            component: Log
        },
        {
            path: "/LogList",
            name: "LogList",
            component: LogList
        }
    ]
});