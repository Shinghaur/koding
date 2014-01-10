# Wrapper for pushing events to Rollbar
KD.logToExternal = (args) ->
  return unless KD.config.logToExternal

  _rollbar?.push args  unless KD.isGuest()

# log ping times so we know if failure was due to user's slow
# internet or our internals timing out
KD.logToExternalWithTime = (name, options)->
  KD.troubleshoot (times)->
    KD.logToExternal {
      options
      msg:"#{name} timed out"
      pings:times
    }

# set user info in rollbar for people tracking
KD.getSingleton('mainController').on "AccountChanged", (account) ->
  return  if KD.isGuest()
  return  unless KD.config.logToExternal and _rollbarParams

  user = KD.whoami?().profile
  _rollbarParams.person =
    id       : user.nickname or "unknown"
    name     : KD.utils.getFullnameFromAccount()
    username : user.nickname
