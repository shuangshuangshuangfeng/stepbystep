## 组件 component

:balloon: 组件是可服用的Vue实例，且带有一个名字。可以通过以下代码创建：

```js
new Vue({el: '#component-demo'}) // el 表示挂载到id为component-demo的组件
```

---------

（1）一个组件的data选项必须是一个函数，因此每个实例可以维护一份**被返回对象的独立的拷贝**。

![](../../assets/2022-11-09-15-21-13-image.png)

（2）一个应用通常会以一棵**嵌套的组件树**的形式来组织。

![](../../assets/2022-11-09-15-21-59-image.png)

:crystal_ball: 组件需要注册后使用，有两种注册类型：

- 全局注册: 使用 `Vue.component`

- 局部注册: **???**

**全局注册代码：**

```js
Vue.component('my-component-name', {
  // ... options ...
})
```







https://v2.cn.vuejs.org/v2/guide/components.html


