import adapter from '@sveltejs/adapter-node';

export default {
  kit: {
    // ... outras configurações
    target: '#svelte',
    adapter: adapter({ 
      // Configurações do adaptador
      port: 8081, // Porta desejada
    }),
  },
};
