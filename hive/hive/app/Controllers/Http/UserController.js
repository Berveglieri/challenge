'use strict'
const User = use('App/Models/User')

class UserController {
  async login ({ view }) {

    return view.render('login')
  }

  async auth ({ auth, request }) {
    const {email, password} = request.all()
    await auth.attempt(email, password)

    return 'Hello World!'
  }

  show ({ auth, params }) {
    if (auth.user.id !== Number(params.id)) {
      return "You cannot see someone else's profile"
    }
    return auth.user
  }

  async register ({ auth, request, view }) {

    return view.render('register')
  }

  async store({ request, response }) {
    const {email, username, password} = request.all()

    await User.create({ email, username, password})

    return response.redirect('/')
  }
}

module.exports = UserController
