const userRoutes = [
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/register/registerView.vue'),
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/loginView.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true,
    },
    component: () => import('@/views/profile/profileView.vue'),
  },
];

export default userRoutes;
