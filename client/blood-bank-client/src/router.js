import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

function loadView(view) {
  return () =>
    import(/* webpackChunkName: "view-[request]" */ `@/views/${view}.vue`)
}
export default new Router({
  routes: [
    {
        path: '/',
        name: 'Home',
        component: loadView('Home'),
        meta: {
          title: 'Home | Blood Bank'
        }
      },
      {
        path: '/donor-form',
        name: 'DonorForm',
        component: loadView('DonorForm'),
        meta: {
          title: 'DonorForm | Blood Bank'
        }
      },
      {
        path: '/edit-donor-form/:formNo',
        name: 'EditDonorForm',
        component: loadView('EditDonorForm'),
        meta: {
          title: 'Edit DonorForm | Blood Bank'
        }
      },
      {
        path: '/image-upload/:formNo',
        name: 'ImageUpload',
        component: loadView('ImageUpload'),
        meta: {
          title: 'ImageUpload | Blood Bank'
        }
      }


  ]
})