import Logo from '../../components/Logo'

export default {
  title: 'Logo',
  component: Logo
}
export const app = () => ({
  render: (h) => h(Logo)
})
app.story = { name: 'Logo' }
