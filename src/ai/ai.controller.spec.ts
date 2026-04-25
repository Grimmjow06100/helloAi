/// <reference types="jest" />
import { Test, TestingModule } from '@nestjs/testing';
import { AiController } from './ai.controller';
import { AiService } from './ai.service';

describe('AiController', () => {
  let controller: AiController;
  let service: AiService;

  const mockAiService = {
    helloAi: jest.fn().mockResolvedValue('I am doing good and you?'),
  };
  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [AiController],
      providers: [
        {
          provide: AiService,
          useValue: mockAiService,
        },
      ],
    }).compile();

    controller = module.get<AiController>(AiController);
    service = module.get<AiService>(AiService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  it('devrait appeler helloAi du service et retourner la réponse', async () => {
    const result = await controller.helloAi();

    expect(result).toBe('I am doing good and you?');
    expect(mockAiService.helloAi).toHaveBeenCalled();
  });
});
