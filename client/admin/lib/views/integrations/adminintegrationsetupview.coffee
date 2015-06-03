kd           = require 'kd'
JView        = require 'app/jview'
remote       = require('app/remote').getInstance()
KDSelectBox  = kd.SelectBox
KDButtonView = kd.ButtonView

# this should be provided by the integrations backend
# putting it here until the backend is ready. - SY
DUMMY_DATA   =
  name    : 'Airbrake'
  logo    : 'https://koding-cdn.s3.amazonaws.com/temp-images/airbrake.png'
  summary : 'Error monitoring and handling.'
  desc    : """
    <p>Airbrake captures errors generated by your team's web applications, and aggregates the results for easy review.</p>
    <p>This integration will post notifications in Slack when an error is reported in Airbrake.</p>
  """
  channels : [
    { name: '#general',   id: '123' }
    { name: '#random',    id: '124' }
    { name: '#off-topic', id: '125' }
  ]
  selectedChannel : '123'
  webhookUrl      : 'https://hooks.koding.com/services/123/abc/def'
  instructions    : """
    ### Setup Instructions

    Here are the steps necessary to add the Airbrake integration.

    ---

    #### Step 1

    In your Airbrake dashboard, click on the menu button in the top navigation, and then select **Integrations** from the sub-menu.

    ![airbrake_step1.png](https://koding-cdn.s3.amazonaws.com/temp-images/airbrake_step1.png)


    #### Step 2

    Click on Slack in the list of integrations, add `https://hooks.slack.com/services/123/abc/def` as the Slack webhook, and choose your notification settings. Press the Save button when you're done.

    ![airbrake_step2.png](https://koding-cdn.s3.amazonaws.com/temp-images/airbrake_step2.png)
  """


module.exports = class AdminIntegrationSetupView extends JView

  constructor: (options = {}, data) ->

    options.cssClass = 'integration-channels'

    super options, data

    selectOptions = []

    for channel in data.channels
      selectOptions.push title: channel.name, value: channel.id

    @channelSelect = new KDSelectBox { selectOptions }

    @addButton = new KDButtonView
      title    : "Add #{data.name} Integration"
      cssClass : 'solid green compact add'
      callback : @bound 'setIntegration'
      loader   : yes

    @cancelButton = new KDButtonView
      title    : 'Cancel'
      cssClass : 'solid red compact cancel'
      callback : @bound 'destroy'


  setIntegration: ->
  
    # assuming we make a call to backend here - SY
    kd.utils.wait 1000, =>

      @destroy()
      @emit 'NewIntegrationAdded', DUMMY_DATA


  pistachio: ->

    { name, desc, summary, logo } = @getData()

    return """
      <header class="integration-view">
        <img src="#{logo}" />
        {p{ #(name)}}
        {{ #(summary)}}
      </header>
      {section.description{ #(desc)}}
      <section class="setup">
        <h4>Post to Channel</h4>
        <label>Start by choosing a channel where exceptions will be posted.</label>
        {{> @channelSelect}}
      </section>
      <section class="buttons">
        {{> @cancelButton}}
        {{> @addButton}}
      </section>
    """
