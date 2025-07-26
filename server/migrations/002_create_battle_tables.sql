-- Create characters table
CREATE TABLE IF NOT EXISTS characters (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(30) NOT NULL,
    level INTEGER NOT NULL DEFAULT 1,
    experience INTEGER NOT NULL DEFAULT 0,
    health INTEGER NOT NULL DEFAULT 100,
    max_health INTEGER NOT NULL DEFAULT 100,
    mana INTEGER NOT NULL DEFAULT 50,
    max_mana INTEGER NOT NULL DEFAULT 50,
    attack INTEGER NOT NULL DEFAULT 10,
    defense INTEGER NOT NULL DEFAULT 5,
    speed INTEGER NOT NULL DEFAULT 10,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create skills table
CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    type VARCHAR(20) NOT NULL, -- "attack", "heal", "buff", "debuff"
    damage INTEGER NOT NULL DEFAULT 0,
    mana_cost INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create character_skills table
CREATE TABLE IF NOT EXISTS character_skills (
    id SERIAL PRIMARY KEY,
    character_id INTEGER NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    skill_id INTEGER NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    level INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create skill_effects table
CREATE TABLE IF NOT EXISTS skill_effects (
    id SERIAL PRIMARY KEY,
    skill_id INTEGER NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    effect_type VARCHAR(20) NOT NULL, -- "damage", "heal", "buff", "debuff"
    value INTEGER NOT NULL,
    duration INTEGER NOT NULL DEFAULT 0, -- 持续回合数
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create battles table
CREATE TABLE IF NOT EXISTS battles (
    id SERIAL PRIMARY KEY,
    player1_id INTEGER NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    player2_id INTEGER NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    winner_id INTEGER REFERENCES characters(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- "pending", "in_progress", "completed"
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create battle_participants table
CREATE TABLE IF NOT EXISTS battle_participants (
    id SERIAL PRIMARY KEY,
    battle_id INTEGER NOT NULL REFERENCES battles(id) ON DELETE CASCADE,
    character_id INTEGER NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    health INTEGER NOT NULL,
    max_health INTEGER NOT NULL,
    mana INTEGER NOT NULL,
    max_mana INTEGER NOT NULL,
    attack INTEGER NOT NULL,
    defense INTEGER NOT NULL,
    speed INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create battle_turns table
CREATE TABLE IF NOT EXISTS battle_turns (
    id SERIAL PRIMARY KEY,
    battle_id INTEGER NOT NULL REFERENCES battles(id) ON DELETE CASCADE,
    turn_number INTEGER NOT NULL,
    actor_id INTEGER NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    action VARCHAR(20) NOT NULL, -- "attack", "skill", "defend", "item"
    target_id INTEGER REFERENCES characters(id) ON DELETE CASCADE,
    result TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_characters_user_id ON characters(user_id);
CREATE INDEX IF NOT EXISTS idx_character_skills_character_id ON character_skills(character_id);
CREATE INDEX IF NOT EXISTS idx_character_skills_skill_id ON character_skills(skill_id);
CREATE INDEX IF NOT EXISTS idx_skill_effects_skill_id ON skill_effects(skill_id);
CREATE INDEX IF NOT EXISTS idx_battles_player1_id ON battles(player1_id);
CREATE INDEX IF NOT EXISTS idx_battles_player2_id ON battles(player2_id);
CREATE INDEX IF NOT EXISTS idx_battle_participants_battle_id ON battle_participants(battle_id);
CREATE INDEX IF NOT EXISTS idx_battle_turns_battle_id ON battle_turns(battle_id);
