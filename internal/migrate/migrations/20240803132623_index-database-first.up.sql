CREATE INDEX idx_polls_user_id ON polls(user_id);
CREATE INDEX idx_options_poll_id ON options(poll_id);
CREATE INDEX idx_votes_user_id ON votes(user_id);
CREATE INDEX idx_votes_option_id ON votes(option_id);