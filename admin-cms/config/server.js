module.exports = ({ env }) => ({
  host: env('HOST', '0.0.0.0'),
  port: env.int('PORT', 1337),
  url: 'http://admin.localhost',
  admin: {
    auth: {
      secret: env('ADMIN_JWT_SECRET', 'd9c0b0e3a406867fb8803bee1393625a'),
    },
  },
});
