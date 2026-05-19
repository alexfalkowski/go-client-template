# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

module Example
  class << self
    def client_config
      @client_config ||= Nonnative::ConfigurationFile.load('.config/client.yml')
    end
  end
end
