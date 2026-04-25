import { Controller, Get } from '@nestjs/common';
import { AiService } from './ai.service';

@Controller('ai')
export class AiController {
  constructor(private readonly aiService: AiService) {}

  @Get('hello')
  public async helloAi() {
    try {
      const response = await this.aiService.helloAi();
      return response;
    } catch {
      console.warn('Error , something went wrong.');
    }
  }
}
