const getters = {
  accessToken: (state) => state.user.accessToken,
  refreshToken: (state) => state.user.refreshToken,
  hasUserInfo: (state) => {
    return JSON.stringify(state.user.userInfo) !== '{}'
  },
  userInfo: (state) => state.user.userInfo
}

export default getters
