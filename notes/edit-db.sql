\c postgres

DO $$ 
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.columns
    WHERE table_name = 'messages'
      AND column_name = 'message_value'
  ) THEN
    ALTER TABLE messages
    ADD COLUMN message_value jsonb;
  END IF;
END $$;