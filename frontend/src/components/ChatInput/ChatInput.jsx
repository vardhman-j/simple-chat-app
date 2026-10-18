import React, { Component } from "react";
import "./ChatInput.scss";

class ChatInput extends Component {
  render() {
    return (
      <div className="ChatInput">
        <label htmlFor="chat-message">Message</label>
        <input
          id="chat-message"
          type="text"
          placeholder="Write a message..."
          onKeyDown={this.props.send}
        />
      </div>
    );
  }
}

export default ChatInput;