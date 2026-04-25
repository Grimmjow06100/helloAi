import { Module } from '@nestjs/common';
import { AiService } from './ai.service';
import { AiController } from './ai.controller';
import { ConfigService } from '@nestjs/config';
import OpenAI from 'openai';

@Module({
  controllers: [AiController],
  providers: [
    {
      provide: 'OPEN_AI_SERVICE',
      useFactory: (configService: ConfigService) => {
        const token = configService.get<string>('OPENAI_API_KEY');
        return new OpenAI({
          apiKey: token,
        });
      },
      inject: [ConfigService],
    },
    AiService,
  ],
})
export class AiModule {}
